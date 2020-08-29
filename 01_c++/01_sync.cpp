#include <iostream>
#include <mutex>
#include <thread>

using namespace std;

mutex m;

void drive(string direction){
    m.lock();
    if(direction == "left"){
        printf("carro entrando na estrada pra ir para esquerda\n");
        printf("carro indo para esquerda\n");
    }
    else if(direction == "right"){
        printf("carro entrando na estrada pra ir para direita\n");
        printf("carro indo para direita\n");
    }
    else{
        printf("Direcao Invalida\n");
    }
    m.unlock();
}

int main(){
    thread t1(drive, "left");
    thread t2(drive, "right");

    t1.join();
    t2.join();
}
