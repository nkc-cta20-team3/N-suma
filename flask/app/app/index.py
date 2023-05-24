# -*- coding: utf-8 -*-
import os
from flask import Flask
from app import static

app = Flask(__name__)
app.register_blueprint(static.app)

@app.route('/')
def index():
  return static.app.send_static_file('index.html')

@app.route("/hello")
def hello():
  return "<h1>Hello, Flask!</h1>"