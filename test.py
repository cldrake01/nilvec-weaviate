import weaviate

# Connect to the local instance of Weaviate
client = weaviate.Client("http://localhost:8080")

# Define a schema for your metadata
def define_schema() -> None:
    schema = {
        "classes": [
            {
                "class": "Book",
                "description": "A class representing a book",
                "properties": [
                    {
                        "name": "title",
                        "dataType": ["string"],
                        "description": "The title of the book"
                    },
                    {
                        "name": "author",
                        "dataType": ["string"],
                        "description": "The author of the book"
                    },
                    {
                        "name": "publishedYear",
                        "dataType": ["int"],
                        "description": "The year the book was published"
                    }
                ]
            }
        ]
    }

    # Add the schema to Weaviate
    try:
        client.schema.create(schema)
        print("Schema created successfully.")
    except Exception as e:
        print(f"Error creating schema: {e}")

# Add metadata to Weaviate
def add_metadata() -> None:
    objects = [
        {
            "class": "Book",
            "properties": {
                "title": "1984",
                "author": "George Orwell",
                "publishedYear": 1949
            }
        },
        {
            "class": "Book",
            "properties": {
                "title": "To Kill a Mockingbird",
                "author": "Harper Lee",
                "publishedYear": 1960
            }
        }
    ]

    # Upload objects to Weaviate
    for obj in objects:
        try:
            client.data_object.create(obj)
            print(f"Uploaded: {obj['properties']['title']}")
        except Exception as e:
            print(f"Error uploading {obj['properties']['title']}: {e}")

if __name__ == "__main__":
    # Ensure schema is defined
    define_schema()

    # Upload metadata
    add_metadata()
