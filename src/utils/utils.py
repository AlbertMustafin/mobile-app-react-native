import os
import re
import datetime
import logging
from typing import Tuple

# Set up logging
logging.basicConfig(level=logging.INFO)

def get_current_time() -> datetime.datetime:
    """Get the current date and time."""
    return datetime.datetime.now()

def validate_email(email: str) -> bool:
    """Validate an email address."""
    pattern = r"^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$"
    return bool(re.fullmatch(pattern, email))

def get_file_extension(filename: str) -> str:
    """Get the extension of a file."""
    return os.path.splitext(filename)[1]

def is_valid_filename(filename: str) -> bool:
    """Check if the filename is valid."""
    valid_chars = set("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._ ")
    return all(char in valid_chars for char in filename)

def get_user_agent(user_agent: str) -> Tuple[str, str, str]:
    """Extract user agent information."""
    pattern = r"Mozilla/5\.0 \(\s*compatible;\s*MSIE\s+(\d+)\.\d+;\s*Windows\s*NT\s+(\d+\.\d+);\s*(.*)"
    match = re.search(pattern, user_agent)
    return match.group(1), match.group(2), match.group(3) if match else ("", "", "")

def get_browser_version(user_agent: str) -> str:
    """Extract browser version from user agent."""
    pattern = r"Mozilla/5\.0 \(\s*compatible;\s*MSIE\s+(\d+)\.\d+;\s*Windows\s*NT\s+\d+\.\d+;\s*(.*)"
    match = re.search(pattern, user_agent)
    return match.group(1) if match else ""

def get_os_version(user_agent: str) -> str:
    """Extract OS version from user agent."""
    pattern = r"Windows\s*NT\s+(\d+\.\d+)"
    match = re.search(pattern, user_agent)
    return match.group(1) if match else ""