base_string = '{\\\\\\\"asset_positions\\\\\\\": [{\\\\\\\"asset_id\\\\\\\": \\\\\\\"0\\\\\\\", \\\\\\\"index\\\\\\\": \\\\\\\"0\\\\\\\", \\\\\\\"quantums\\\\\\\": \\\\\\\"100000000000000000\\\\\\\"}], \\\\\\\"id\\\\\\\": {\\\\\\\"number\\\\\\\": \\\\\\\"XXX\\\\\\\", \\\\\\\"owner\\\\\\\": \\\\\\\"dydx1fjg6zp6vv8t9wvy4lps03r5l4g7tkjw9wvmh70\\\\\\\"}, \\\\\\\"margin_enabled\\\\\\\": true}'

result = ''
for i in range(10000):
    result += base_string.replace('XXX', str(i))

print(result)