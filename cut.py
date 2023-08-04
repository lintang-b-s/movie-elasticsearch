import json
file_name = "movies-tenflix-2.json"

with open(file_name, "r", encoding='utf-8') as f:
    my_list = json.load(f)

    new_data = []

    for idx, obj in enumerate(my_list):
        if int(obj["releaseYear"]) > 2010:
            obj["releaseYear"] = int(obj["releaseYear"])
            obj["cast"] = obj["cast"].split(',')
            obj["genre"] = obj["genre"].split(',')
            new_data.append(obj)

new_file_name = "movies-tenflix.json"

with open(new_file_name, 'w', encoding='utf-8') as f:
      json.dump(new_data, f, indent=4)