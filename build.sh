#! /bin/bash
# go install -o sre
# git describe --always --long
BUILD_FLAG=true
FIX_VERSION=0.0.1
GIT_DESCRIBE=`git describe --tags --always`
# echo "Version : $FIX_VERSION  $GIT_DESCRIBE"
# increment the build number
VERSION=`echo $GIT_DESCRIBE | awk '{split($0,a,"-"); print a[1]}'`
BUILD=`echo $GIT_DESCRIBE | awk '{split($0,a,"-"); print a[2]}'`
PATCH=`echo $GIT_DESCRIBE | awk '{split($0,a,"-"); print a[3]}'`

if [[ "${GIT_DESCRIBE}" =~ ^[A-Fa-f0-9]+$ ]]; then
    VERSION="0.0.0"
    BUILD=`git rev-list HEAD --count`
    PATCH=${GIT_DESCRIBE}
fi

if [ "${BUILD}" = "" ]; then
    BUILD='0'
fi

if [ "${BUILD}" = "" ]; then
    PATCH=$GIT_DESCRIBE
fi
echo ${VERSION}+build.${BUILD}.${PATCH}

if [[ "$BUILD_FLAG" == true  ]] ; then
  # Proceed with the build if BUILD=true
  echo "✅ - Build can proceed ⏳"
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
#   exit 1;
else
 exit 1;
fi
# Done!
echo
echo "==> Results:"
ls -hl /usr/local/bin/sre
# sre health
