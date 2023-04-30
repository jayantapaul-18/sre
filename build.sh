# #!/bin/sh
# go install -o sre
# git describe --always --long
VERSION=0.0.1
echo "Version : $VERSION"
# echo $VERSION
echo "******************** build **************************************************"
go build -o sre -ldflags="-X 'main.Version=v0.0.1'"
echo "******************** copy config to $HOME *****************************"
cp sre-config.json $HOME/sre-config.json
chmod 777 $HOME/sre-config.json
echo "******************** move pkg to bin ********************"
mv sre /usr/local/bin/
echo "******************** change permission to execute sre cli********************"
chmod +x /usr/local/bin/sre
echo "******************** ready to use sre cli ***********************************"
# sre health
