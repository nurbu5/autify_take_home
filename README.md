# Some notes
- The image count is simply checking for "img" tags, so images being pulled in as css b
  - Not sure how accurate the image and link checks are
- If I had more time I probably would have written some specs
- Make sure to include the protocol in the URL or else you'll get an error (Ex: https://google.com)

# Usage
- docker build . --tag fetch
- docker run --volume $(pwd):/app fetch https://google.com
- or docker run fetch --metadata https://google.com
