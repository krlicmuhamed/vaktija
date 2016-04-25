# Make sure only root can run our script
if [ "$(id -u)" != "0" ]; then
   echo "This install script must be run as root." 1>&2
   exit 1
fi
export GOPATH=/
go get github.com/PuerkitoBio/goquery
go install
printf "%s" "/bin/vaktija" >> /usr/bin/vaktija
echo "Done. Run with vaktija."
