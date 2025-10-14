// Timestamp: 2025-10-14 01:06:14

const calculateDelay = (retryCount) => {
    return Math.pow(2, retryCount) * 1000;
};

