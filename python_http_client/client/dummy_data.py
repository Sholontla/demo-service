import uuid
import random
import numpy as np
from faker import Faker
import faker_commerce


class DummyData:
    def __init__(self):
        self.fake = Faker()
        self.fake.add_provider(faker_commerce.Provider)
        self.fake.add_provider(faker_commerce.Provider)
        self.gender = ["M", "F"]

    def dummy_data_creation(self) -> dict:
        producer_id = str(uuid.uuid4())

        # producer_message
        producer_message_id = str(uuid.uuid4())
        host = self.fake.ascii_email()
        client = self.fake.name()
        ip = self.fake.ipv4_private()
        port = np.random.randint(1000, 10000, size=1)

        # data_producer
        data_producer_id = str(uuid.uuid4())

        product = self.fake.ecommerce_name()
        name = self.fake.ecommerce_name()
        category = self.fake.ecommerce_name()
        sub_category = self.fake.ecommerce_name()
        price = np.random.uniform(.10, 1000, 1)
        quantity = np.random.randint(1, 20, size=1)
        supplier = self.fake.company()
        description = self.fake.sentence(nb_words=7)
        gender = random.choice(self.gender)
        information_created_at = str(self.fake.date_this_month())
        producer_service_area = self.fake.address()
        producer_created_at = str(self.fake.date_this_month())

        new_data = {
            "producer_id": producer_id,
            "producer_message": {
                "producer_message_id": producer_message_id,
                "host": host,
                "client": client,
                "ip": ip,
                "port": int(port),
                "data_producer": {
                    "data_producer_id": data_producer_id,
                    "product": product,
                    "name": name,
                    "category": category,
                    "sub_category": sub_category,
                    "price": float(price),
                    "quantity": int(quantity),
                    "supplier": supplier,
                    "description": description,
                    "gender": gender
                },
                "information_created_at": information_created_at
            },
            "producer_service_area": producer_service_area,
            "producer_created_at": producer_created_at
        }

        return new_data
