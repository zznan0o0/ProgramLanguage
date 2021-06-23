import jwt

jwtstr = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjhtMEs4VFRYdDJHQ1dTZlpydTJ1UVEiLCJ0eXAiOiJKV1QifQ.eyJuYmYiOjE2MjMxMzUxMjUsImV4cCI6MTYyMzEzNTQyNSwiaXNzIjoiaHR0cDovLzE5Mi4xNjguMTAuMTUzOjQwMDQiLCJhdWQiOiJVc2VyQ2VudGVyQ29kZSIsImlhdCI6MTYyMzEzNTEyNSwiYXRfaGFzaCI6IjhKTlhYQUxNTi1ZSHhMRTc3T0VSUGciLCJzaWQiOiJ4TTRoOG91LUtoZ2Q4cUZ6WEZuTWdnIiwic3ViIjoiOTk5IiwiYXV0aF90aW1lIjoxNjIzMTM1MTIyLCJpZHAiOiJsb2NhbCIsImFtciI6WyJwd2QiXX0.XqE8NhjX42IvK1anjET541FHmuXaWg60WG8Td7Gtu1YS7lEPHPKElIu1iI-mpqeeBzEF9hVLvUgu1QA13Jx8P17GGANHUo2pPfRDnnYObDYIVuKrJ1Fx25dMKtNzQU6tT76xkV5yezLHCFFXaxStSZybdW90U88xV89v_4cQS5VyEXyYu-n2aIvqs9W3FxaaCPasvyhGuyCs0k7RlgoOt6UV8JuW9LiN1VCiEAQwNG7dKIZgMH07chF-umniX6B5QVDJWhom0-s3NQ9uCFkRMR0fCmeqeCMAaAGMBtQRv_-Pxkn7dSLZR6ojZlqIMAEoKMOWRR5Kd-eiVIx3D-8Hvw"


s = jwt.decode(jwtstr, algorithms=['HS256'], options={'verify_signature': False})
print(s)

# key = "secret"
# token = jwt.encode({"test": "100"},key,"HS256")
# print(token)
# header,payload,signature = token.split(b".")

# import base64

# def equal(b: bytes):
#     """用来补齐被JWT去掉的等号"""
#     rest = len(b) % 4
#     return b + b'=' * rest

# print("header=",base64.urlsafe_b64decode(equal(header)))
# print("payout=", base64.urlsafe_b64decode(equal(payload)))
# print("signature", base64.urlsafe_b64decode(equal(signature)))