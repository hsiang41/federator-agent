#!/bin/sh
#
[ "${AIHOME}" = "" ] && export AIHOME=/opt/alameda/federatorai-agent

#
while :
do
    cd ${AIHOME}/bin
    ${AIHOME}/bin/transmitter run
    [ -f /tmp/.pause ] && sleep 300 || sleep 30
done

exit 0
