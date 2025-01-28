import inspect
import os

import urllib3

import pyaes
from pyaes import AESModeOfOperationCBC, Decrypter


class Api:
    seed = None

    def __init__(self):
        with open("../.session", 'r') as file:
            self.session = file.readline()
            self.headers = {
                'Cookie': f"everybody-codes={self.session}",
            }

    def get_seed(self) -> str:
        if self.seed is None:
            response = urllib3.request(
                'GET',
                "https://everybody.codes/api/user/me",
                headers=self.headers,
            )

            self.seed = response.json()['seed']

        return self.seed

    def get_encrypted_inputs(self, event: int, quest: int) -> [str, str, str]:
        seed = self.get_seed()

        response = urllib3.request(
            'GET',
            f"https://everybody-codes.b-cdn.net/assets/{event}/{quest}/input/{seed}.json",
            headers=self.headers,
        )

        response_json = response.json()
        return response_json['1'], response_json['2'], response_json['3']

    def get_decryption_keys(self, event: int, quest: int) -> [str | None, str | None, str | None]:
        response = urllib3.request(
            'GET',
            f"https://everybody.codes/api/event/{event}/quest/{quest}",
            headers=self.headers,
        )

        response_json = response.json()
        return response_json.get('key1'), response_json.get('key2'), response_json.get('key3')

    def get_inputs(self) -> [str | None, str | None, str | None]:
        calling_file_path_directories = inspect.stack()[1].filename.split(os.path.sep)
        event = int(calling_file_path_directories[-2])
        quest = int(calling_file_path_directories[-1].split('.')[0])

        encrypted_inputs = self.get_encrypted_inputs(event, quest)
        decryption_keys = self.get_decryption_keys(event, quest)
        decrypted_inputs = []

        for encrypted_input, decryption_key in zip(encrypted_inputs, decryption_keys):
            if not decryption_key:
                decrypted_inputs.append(None)
                continue

            decryption_key = f"{decryption_key[:20]}~{decryption_key[21:]}"

            encrypted_input_bytes = bytes.fromhex(encrypted_input)
            decryption_key_bytes = decryption_key.encode('utf-8')
            initialization_vector_bytes = decryption_key[:16].encode('utf-8')

            decrypter = Decrypter(AESModeOfOperationCBC(decryption_key_bytes, iv=initialization_vector_bytes))
            decrypted_input = decrypter.feed(encrypted_input_bytes)
            decrypted_input += decrypter.feed()

            decrypted_inputs.append(decrypted_input.decode('utf-8'))

        return decrypted_inputs[0], decrypted_inputs[1], decrypted_inputs[2]

