#!/bin/sh

youtube-dl -x --exec 'mpv {}' $@
