from flask import Flask, request, jsonify
import pytesseract
from PIL import Image
import re

app = Flask(__name__)

def parse_bill_text(text):
    # Define a list to store parsed data for items
    items = []

    # Updated pattern to capture Nepali and English words, handling optional spaces
    item_pattern = re.compile(
        r'(?P<number>\d+)\s+(?P<name>[\w\s]+?)\s+(?P<quantity>\d+)\s+(?P<rate>\d+)\s+(?P<total>\d+)'
    )

    # Iterate through each line and search for matches to the pattern
    for line in text.splitlines():
        match = item_pattern.search(line)
        if match:
            try:
                # Extract each field with its appropriate type
                number = int(match.group("क्र स"))
                name = match.group("विवरण").strip()
                quantity = int(match.group("मात्रा"))
                rate = int(match.group("दर"))
                total = int(match.group("रकम"))

                # Append the parsed item data to the items list
                items.append({
                    "number": number,
                    "name": name,
                    "quantity": quantity,
                    "rate": rate,
                    "total": total
                })
            except ValueError:
                # If there's a conversion error, skip this line
                print(f"Skipping line due to conversion error: {line}")
                continue

    return items

@app.route('/process-image', methods=['POST'])
def process_image():
    # Get the uploaded image file
    file = request.files['image']
    image = Image.open(file)

    # Perform OCR on the image to extract text
    try:
        text = pytesseract.image_to_string(image, lang="nep")
    except pytesseract.TesseractError as e:
        # Handle potential OCR errors
        return jsonify({"error": str(e)}), 500

    # Parse the text to extract relevant bill information
    items = parse_bill_text(text)
    
    # Create the response in JSON format
    response = {
        "items": items
    }

    return jsonify(response)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
