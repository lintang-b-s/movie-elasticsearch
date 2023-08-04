import json
import requests
import sys


# create mapping for index movieswiki
urlMap  = "http://localhost:9200/movieswiki"
fMap = open('movie-mapping.json')
headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
dataMap = json.load(fMap)
# print(dataMap["peta"])
rMap = requests.put(urlMap, data=json.dumps(dataMap["peta"]), headers=headers)
print(rMap.status_code)

fMap.close()


# indexing movieswiki index with document in movies-tenflix.json
url = "http://localhost:9200/{}/_doc".format("movieswiki")

# Opening JSON file
f = open('movies-tenflix.json')

# returns JSON object as
# a dictionary
data = json.load(f)

# print json data
#print(data['prizes'][0])

# print number of documents in the json
count = len(data)
print(count)

# index each document
for each in range(count):
    movie = data[each]
    r = requests.post(url, data=json.dumps(movie), headers=headers)
    # print(r.status_code)

# Closing file
f.close()