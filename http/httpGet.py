import requests

class LogWriter:
    def write(self, data):
        print(data.decode())
        print("Just wrote this many bytes:", len(data))

def main():
    response = requests.get("http://google.com")
    if response.status_code != 200:
        print("Error:", response.status_code)
        exit(1)

    lw = LogWriter()
    lw.write(response.content)

if __name__ == "__main__":
    main()
