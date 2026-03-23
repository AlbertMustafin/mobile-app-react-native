import os
import re
import datetime
import logging

# Set up logging
logging.basicConfig(level=logging.INFO)

def get_current_time():
    return datetime.datetime.now()

def validate_email(email: str) -> bool:
    """Validate an email address."""
    pattern = r"^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$"
    return bool(re.match(pattern, email))

def get_file_extension(filename: str) -> str:
    """Get the extension of a file."""
    return os.path.splitext(filename)[1]

def is_valid_filename(filename: str) -> bool:
    """Check if the filename is valid."""
    valid_chars = set("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._ ")
    return all(char in valid_chars for char in filename)

def get_user_agent(user_agent: str) -> tuple:
    """Extract user agent information."""
    pattern = r"Mozilla/5\.0 \(compatible; MSIE (\d+)\.0; Windows NT (\d+\.\d+); (.*); .+"
    match = re.search(pattern, user_agent)
    return match.group(1), match.group(2), match.group(3)