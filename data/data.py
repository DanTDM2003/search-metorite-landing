import json
from datetime import datetime, timezone

# Load the JSON data from the file
input_file = 'meteorite-landings.json'
output_file = 'meteorite-landings-api-1_converted.json'

with open(input_file, 'r') as file:
    data = json.load(file)

# Convert the specified fields
for entry in data:
    if 'mass' not in entry:
        entry['mass'] = 0.0
    else:
        entry['mass'] = float(entry['mass'])
    
    if 'reclat' not in entry:
        entry['reclat'] = 0.0
    else:
        entry['reclat'] = float(entry['reclat'])

    if 'reclong' not in entry:
        entry['reclong'] = 0.0
    else:
        entry['reclong'] = float(entry['reclong'])
    entry['id'] = int(entry['id'])

    if 'geolocation' not in entry:
        entry['geolocation'] = {'latitude': 0.0, 'needs_recoding': False, 'longitude': 0.0}
    else:
        entry['geolocation']['latitude'] = float(entry['geolocation']['latitude'])
        entry['geolocation']['longitude'] = float(entry['geolocation']['longitude'])

    dt = datetime.strptime(entry['year'], "%Y-%m-%dT%H:%M:%S").replace(tzinfo=timezone.utc)
    entry['year'] = dt.isoformat()

# Write the modified data back to a new JSON file
with open(output_file, 'w') as file:
    json.dump(data, file, indent=4)

print("Conversion complete. The modified file has been saved as:", output_file)
