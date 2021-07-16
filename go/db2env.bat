@echo off

set PATH=%CD%\clidriver\bin;%PATH%
set DB2DSN=HOSTNAME=localhost;DATABASE=testdb;PORT=50001;PROTOCOL=TCPIP;UID=db2inst1;PWD=pa55w0rd

@echo on
