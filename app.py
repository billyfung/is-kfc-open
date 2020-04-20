from flask import Flask, render_template
from datetime import datetime
import pytz
import os

app = Flask(__name__)

@app.route('/')
def is_open():
    tz_NZ = pytz.timezone('Pacific/Auckland')
    kfc_date = tz_NZ.localize(datetime(2020,4,27, hour=23, minute=59))
    tz_UTC = pytz.timezone('UTC')
    if datetime.utcnow().astimezone(tz_UTC) > kfc_date:
        return render_template('index.html', result=1)
    else:
        return render_template('index.html', result=0)


if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0',port=int(os.environ.get('PORT', 8080)))