# utils.py

import os
import re
import shutil
import yaml

def load_yaml(file_path):
    """Loads a YAML file into a dictionary."""
    with open(file_path, 'r') as file:
        return yaml.safe_load(file)

def get_file_contents(file_path):
    """Returns the contents of a file as a string."""
    with open(file_path, 'r') as file:
        return file.read()

def get_file_hash(file_path):
    """Returns the SHA-256 hash of a file."""
    with open(file_path, 'rb') as file:
        return hashlib.sha256(file.read()).hexdigest()

def copy_file(src, dst):
    """Copies a file from src to dst."""
    shutil.copy2(src, dst)

def create_dir(dir_path):
    """Creates a directory at dir_path."""
    os.makedirs(dir_path, exist_ok=True)

def remove_dir(dir_path):
    """Removes a directory at dir_path."""
    shutil.rmtree(dir_path)

def is_dir_empty(dir_path):
    """Returns True if dir_path is empty, False otherwise."""
    return len(os.listdir(dir_path)) == 0

def is_file_locked(file_path):
    """Returns True if file_path is locked, False otherwise."""
    fd = os.open(file_path, os.O_RDONLY)
    try:
        os.flock(fd, os.LOCK_EX | os.LOCK_NB)
    except BlockingIOError:
        os.close(fd)
        return True
    os.close(fd)
    return False

def is_port_in_use(port):
    """Returns True if port is in use, False otherwise."""
    import socket
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        return s.connect_ex(('localhost', port)) == 0