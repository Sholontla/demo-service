import time
import json
import requests
from dummy_data import DummyData


class HttpClient:
    def __init__(self):
        self.url_post1 = "http://localhost:8004/service/producer/register1"
        self.url_post2 = "http://localhost:8004/service/producer/register2"
        self.url_post3 = "http://localhost:8004/service/producer/register3"
        self.url_post4 = "http://localhost:8004/service/producer/register4"
        self.url_post5 = "http://localhost:8004/service/producer/register5"

    def client1(self, new_data1):

        post_response = requests.post(self.url_post1, json=new_data1)
        return post_response.json()

    def client2(self, new_data2):
        post_response = requests.post(self.url_post2, json=new_data2)
        return post_response.json()

    def client3(self, new_data3):
        post_response = requests.post(self.url_post3, json=new_data3)
        return post_response.json()

    def client4(self, new_data4):
        post_response = requests.post(self.url_post4, json=new_data4)
        return post_response.json()

    def client5(self, new_data5):
        post_response = requests.post(self.url_post5, json=new_data5)
        return post_response.json()


class SchedulerRequest:
    def __init__(self):
        self.time = starttime = time.time()

        self.dummy = DummyData()
        self.http = HttpClient()

    def request(self):

        while True:
            self.http.client1(self.dummy.dummy_data_creation())
            time.sleep(1.0 - ((time.time() - self.time) % 1.0) / 1.0)

            self.http.client2(self.dummy.dummy_data_creation())
            time.sleep(1.0 - ((time.time() - self.time) % 1.0) / 1.0)

            self.http.client3(self.dummy.dummy_data_creation())
            time.sleep(1.0 - ((time.time() - self.time) % 1.0) / 1.0)

            self.http.client4(self.dummy.dummy_data_creation())
            time.sleep(1.0 - ((time.time() - self.time) % 1.0) / 1.0)

            self.http.client5(self.dummy.dummy_data_creation())
            time.sleep(1.0 - ((time.time() - self.time) % 1.0) / 1.0)


if __name__ == "__main__":
    scheduler = SchedulerRequest()

    scheduler.request()
