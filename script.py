from kafka import KafkaProducer
from kafka.errors import KafkaError

producer = KafkaProducer(bootstrap_servers=['127.0.0.1:29092', '127.0.0.1:39092', '127.0.0.1:49092'])

future = producer.send('driver-in', b'''{
    "id": "284655d6-0190-49e7-34e9-9b4060acc261",
    "source": "/trip",
    "type": "trip.event.created",
    "datacontenttype": "application/json",
    "time": "2023-11-09T17:31:00Z",
    "data": {
        "trip_id": "e82c42d6-b86f-4e2a-93a2-858413acb148",
        "offer_id": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InN0cmluZyIsImZyb20iOnsibGF0IjowLCJsbmciOjB9LCJ0byI6eyJsYXQiOjAsImxuZyI6MH0sImNsaWVudF9pZCI6InN0cmluZyIsInByaWNlIjp7ImFtb3VudCI6OTkuOTUsImN1cnJlbmN5IjoiUlVCIn19.fg0Bv2ONjT4r8OgFqJ2tpv67ar7pUih2LhDRCRhWW3c"
    }
}''')

# Block for 'synchronous' sends
try:
    record_metadata = future.get(timeout=10)
except KafkaError as err:
    # Decide what to do if produce request failed...
    print(err)
    pass