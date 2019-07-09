#!/bin/sh
#
[ "${AIHOME}" = "" ] && export AIHOME=/opt/alameda/federatorai-agent

# create pwent entry because logrotate need the running uid inside /etc/passwd
MY_UID="`id -u`"
if [ "`cat /etc/passwd | awk -F: '{print $3}' | grep \"^${MY_UID}$\"`" = "" ]; then
    sed -e "s/alameda:/alameda2:/g" /etc/passwd > /tmp/.tmpfn; cat /tmp/.tmpfn > /etc/passwd; rm -f /tmp/.tmpfn
    echo "alameda:x:${MY_UID}:0:Federator.ai:${AIHOME}:/bin/sh" >> /etc/passwd
fi

do_crond()
{
    sleep_time=3600
    while :; do
        case "`date +%H`" in
            "00") # hour is 00
                logrotate -v -f /etc/logrotate.conf
                ;;
        esac
        sleep ${sleep_time}
    done
    exit 0
}

do_liveness()
{
    # ${AIHOME}/bin/transmitter probe | grep ", READY$" > /dev/null
    # [ "$?" != "0" ] && echo "Failed in server liveness probe." && return $?
    return 0
}

do_readiness()
{
    # ${AIHOME}/bin/transmitter probe | grep ", READY$" > /dev/null
    # [ "$?" != "0" ] && echo "Failed in server readiness probe." && return $?
    return 0
}

do_start()
{
    # start crond
    $0 crond &

    # start main service
    while :; do
        cd ${AIHOME}/bin
        ${AIHOME}/bin/transmitter run
        [ -f /tmp/.pause ] && sleep 300 || sleep 30
    done
    return $?
}

show_usage()
{
    /bin/echo -e "\n\nUsage: $0 [crond|liveness|readiness|start]\n\n"
    exit 1
}

#
# Main
#

# start crond only
case "$1" in
    "crond")
        do_crond
        exit $?
        ;;
    "liveness")
        do_liveness
        exit $?
        ;;
    "readiness")
        do_readiness
        exit $?
        ;;
    "start")
        do_start
        ;;
    *)
        show_usage
        exit $?
        ;;
esac

exit 0
