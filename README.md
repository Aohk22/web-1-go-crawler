**To dos**

- [x] Functional.
- [ ] Implement concurrency or something that can make use of many queues.
- [ ] Save HTML files.

# Description

For learning Golang.

References:
- https://bytebytego.com/courses/system-design-interview/design-a-web-crawler
- https://en.wikipedia.org/wiki/Web_crawler

# Demonstration

## First iteration

The seed URL is put during init.  
Url Downloader gets the page -> extract list of URLs -> pass into Url Frontier (retained parent URl for parsing relative links).  
Url Frontier then passes processed URLs into queue.

A queue is selected by URL Downloader using round robin scheduling for simplicity.

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/71e3a29b-3469-4792-b786-f144d94562e2" />

## Second iteration

Here is there extracted URLs, they are mapped to queues using hostnames.  
As there are many relative paths in this page, it might be beneficial to implement a weighted queue selection.

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/7a59661c-3c0c-4e48-88e2-4bad7a69565f" />

### Relative paths do work

<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/c1a1fba1-9c9e-4312-8bd3-14a15f35bc6c" />
