# -*- coding: utf-8 -*-
from flask import Blueprint

app = Blueprint('public', __name__, static_url_path='', static_folder='../public')