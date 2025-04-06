### RingBuffer进程间通信

writer.cpp

```c++
#include <iostream>
#include <sys/mman.h>
#include <fcntl.h>
#include <unistd.h>
#include <cstring>
#include <chrono>
#include <thread>

const char* shm_name = "/my_shm";

struct MyData {
    int number;
    char str[100];
};

struct RingBuffer {
    MyData buffer[5];
    size_t head;
    size_t tail;
    bool full;
};

void push(RingBuffer& rb, const MyData& value) {
    rb.buffer[rb.head] = value;
    rb.head = (rb.head + 1) % 5;
    if (rb.full) {
        rb.tail = (rb.tail + 1) % 5;
    }
    rb.full = (rb.head == rb.tail);
}

int main() {
    int fd = shm_open(shm_name, O_CREAT | O_RDWR, 0666);
    ftruncate(fd, sizeof(RingBuffer));
    RingBuffer* rb = (RingBuffer*)mmap(NULL, sizeof(RingBuffer), PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    close(fd);
    std::memset(rb, 0, sizeof(RingBuffer));

    int data_value = 0;
    while (true) {
        MyData data = {data_value++, "example_string"};
        push(*rb, data);
        std::cout << "写入数据: " << data.number << ", " << data.str << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    return 0;
}

```

reader.cpp

```c++
#include <iostream>
#include <sys/mman.h>
#include <fcntl.h>
#include <unistd.h>
#include <cstring>
#include <semaphore.h>
#include <chrono>
#include <thread>

const char* shm_name = "/my_shm";
const char* sem_read_name = "/sem_read";

struct MyData {
    int number;
    char str[100];
};

struct RingBuffer {
    MyData buffer[5];
    size_t head;
    size_t tail;
    bool full;
};

MyData pop(RingBuffer& rb) {
    if (rb.head == rb.tail && !rb.full) {
        return MyData{-1, ""};
    }
    MyData value = rb.buffer[rb.tail];
    rb.tail = (rb.tail + 1) % 5;
    rb.full = false;
    return value;
}

int main() {
    int fd = shm_open(shm_name, O_RDWR, 0666);
    RingBuffer* rb = (RingBuffer*)mmap(NULL, sizeof(RingBuffer), PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    close(fd);

    sem_t* sem_read = sem_open(sem_read_name, O_CREAT, 0666, 0);

    while (true) {
        if (rb->head != rb->tail || rb->full) {
            MyData data = pop(*rb);
            if (data.number != -1) {
                std::cout << "读取数据: " << data.number << ", " << data.str << std::endl;
            }
            sem_post(sem_read);
        } else {
            std::this_thread::sleep_for(std::chrono::seconds(1));  // 等待数据写入
        }
    }

    sem_close(sem_read);
    return 0;
}

```

