FROM python:3.11.3

WORKDIR /app

ADD requirements.txt /app
RUN pip install --upgrade pip && pip install -r requirements.txt 

CMD ["python", "index.py"]