import subprocess
import sys


# Function to install a package using pip
def install_package(package):
    subprocess.check_call([sys.executable, "-m", "pip", "install", package])


# Check if requests is installed, if not, install it
try:
    import requests
except ImportError:
    install_package("requests")
    import requests

from datetime import datetime
import random

# Define your data
user_ids = ["user1", "user2", "user3", "user4", "user5"]
products = ["product1", "product2", "product3", "product4", "product5"]
actions = ["view", "click", "add_to_cart"]


# Function to get a formatted timestamp
def get_formatted_timestamp():
    return datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%SZ")


# Send 20 random requests
for _ in range(500):
    # Select random user_id, product, and action
    user_id = random.choice(user_ids)
    product = random.choice(products)
    action = random.choice(actions)

    # Create user interaction
    user_interaction = {
        "user_id": user_id,
        "product_sku": product,
        "action": action,
        "interaction_timestamp": get_formatted_timestamp(),
        "interaction_duration": random.randint(1, 100)
    }

    # Send POST request
    response = requests.post("http://localhost:8080/user_interaction", json={"user_interaction": [user_interaction]})
    print(response.text)
