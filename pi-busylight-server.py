#!/usr/bin/env python

import signal
import time
import scrollphathd
from flask import Flask
import logging

app = Flask(__name__)

#Disable logging
log = logging.getLogger('werkzeug')
log.disabled = True
app.logger.disabled = True


@app.route('/')
def home():
    return "pi-busylight is running"

@app.route('/api/on', methods=['POST'])
def busylight_on():
    scrollphathd.fill(brightness=0.5)
    scrollphathd.show()
    return "on"

@app.route('/api/off', methods=['POST'])
def busylight_off():
    scrollphathd.clear()
    scrollphathd.show()
    return "off"

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)