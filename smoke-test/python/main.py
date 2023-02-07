import socket
import os
import _thread

HOST, PORT = "0.0.0.0", 8082


def send_echo(conn):
    while True:
        data = conn.recv(1024)
        if not data:
            break
        conn.sendall(data)
    conn.close()


if __name__ == "__main__":

    my_socket = socket.socket()
    my_socket.bind((HOST, PORT))
    my_socket.listen(10)

    print("listening...")
    try:
        while True:
            conn, addr = my_socket.accept()
            print(f"received new connection from: {addr}")
            _thread.start_new_thread(send_echo, (conn,))

    except KeyboardInterrupt:
        print("exiting...")
    except Exception as e:
        print("something went wrong: ", e.with_traceback())
    finally:
        my_socket.close()
