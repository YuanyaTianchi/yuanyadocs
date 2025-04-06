import requests
from pprint import pprint

def get_location_info(latitude, longitude):
    ak = "UGzhk1h30O5IdAoDM7KZHcOL5qjTcSY2"
    url = "http://api.map.baidu.com/reverse_geocoding/v3/"
    params = {
        "ak": ak,
        "location": f"{latitude},{longitude}",
        "output": "json"
    }
    response = requests.get(url, timeout=5, params=params)
    data = response.json()
    pprint(data)
    if data["status"] == 0:
        result = data["result"]
        province = result["addressComponent"]["province"]
        city = result["addressComponent"]["city"]
        district = result["addressComponent"]["district"]
        return f"省份：{province}，城市：{city}，区县：{district}"
    else:
        return "抱歉，无法获取到省市区信息。"

# 示例经纬度：北京天安门
latitude = 39.915168
longitude = 116.403875
location_info = get_location_info(latitude, longitude)
print(location_info)