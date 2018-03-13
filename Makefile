###########################################################################################
# WARNING!!! This script is for developping purposes only. Don't try to use it to install #
# application.                                                                            #
###########################################################################################

BUILDDIR := $(CURDIR)/build/amd64/gichidan

CC = go build --ldflags "-X main.VERSION=1.1.1"
TARGET = gichidan
ARTEFACT = cliface

.PHONY: all install	

all: $(TARGET)

$(TARGET): 
	$(CC)

install: $(TARGET)
	cp $(TARGET) $(BUILDDIR)
	rm -f $(TARGET)
	rm -f $(ARTEFACT)
