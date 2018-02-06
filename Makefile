##############################################################
# @Package gichidan                                          #
#                                                            #
# Makefile                                                   #
##############################################################

IMPORT_PATH := github.com/hIMEI29A/gichidan
BUILDDIR := $(CURDIR)/build/amd64/gichidan
GITHUB_REPO := hIMEI29A/gichidan

CC = go build
TARGET = gichidan
ARTEFACT = cliface

.PHONY: all clean install	

all: $(TARGET) clean

clean:
	rm -f $(TARGET)

$(TARGET): 
	$(CC)

install: $(TARGET)
	cp $(TARGET) $(BUILDDIR)
	rm -f $(ARTEFACT)
