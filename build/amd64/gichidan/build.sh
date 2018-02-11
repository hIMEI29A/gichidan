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
   echo "Usage : build.sh -v|--version [VERSION]" 
}

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
        echo "Version = $VERSION"
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
