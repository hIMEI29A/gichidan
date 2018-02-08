#!/usr/bin/env bash

###########################################################################################
# WARNING!!! This script is for developping purposes only. Don't try to use it to install #
# application.                                                                            #
###########################################################################################

# Functions
#########################

# Prints usage message end exit
usage()
{
   echo "Usage : build.sh -v|--version [VERSION] -k|--key [KEY]" 
}

# Uploads .deb to BinTray
#upload()
# {
#    curl -v -X PUT -T "$1" -uhimei29a:"$2" "https://api.bintray.com/content/himei29a/deb/gichidan/"$3"/pool/main/$1;deb_distribution=jessie;deb_component=main;deb_architecture=amd64;publish=1"
#}
#

# Variables 
########################

FILE=*.deb
REMOVE=gichidan

# ClI args processing
########################

# if there is no args provided, print usage and exit
if [ "$1" = "" ]
then
    usage
    exit
fi

# args parsing
while [ "$1" != "" ]; do
    case $1 in
        -v | --version )
        VERSION="$2"
        shift
        shift
        echo "$VERSION"
        ;;
        -k | --key )
        KEY="$2"
        shift
        shift
        echo "$KEY"
        ;;
        * )
        usage
        exit 1
    esac
done

# Package building 
#######################

mkdeb build -version="$VERSION" mkdeb.json
rm "$REMOVE"
echo "Binary removed"

# Upload package 
#######################
#
#for f in $FILE
#do
#    echo "Deb package $f found"
#    upload "$f" "$KEY" "$VERSION"
#done
