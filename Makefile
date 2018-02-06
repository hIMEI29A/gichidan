###########################################
#          @Package gichidan              #
#                                         #
#              Makefile                   #
###########################################

BUILDDIR := $(CURDIR)/build/amd64/gichidan

CC = go build --ldflags "-X main.VERSION=1.0.0"
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
	rm -f $(TARGET)
