import requests
import pymysql

db_config = {
    'host': 'ts.dn11.top',
    'port': 33066,
    'user': 'root',
    'password': 'SXFsvsxegJ4PtI84MR',
    'database': 'tiktok_startup'
}

db = pymysql.connect(**db_config)

cursor = db.cursor()
cursor.execute("SELECT username FROM user")
usernames = cursor.fetchall()
cursor.close()
db.close()

base_url = 'http://127.0.0.1:8888/douyin/user/register/'

for username in usernames:

    params = {
        'username': username[0].replace(" ", "☆"),
        'password': 'password'
    }
    response = requests.post(base_url, params=params)

    print(f'{username}: {response.status_code} - {response.text}')
