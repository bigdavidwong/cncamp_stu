1. build Dockerfile in module3 path(include main file)
`docker build -t gohttpserver .`
![img.png](img.png)

2. push to docker hub;

`docker tag gohttpserver:latest bigdavidwong/camp:v1.0`

`docker push bigdavidwong/camp:v1.0`
![img_1.png](img_1.png)

3. docker run container;

`docker run -p 8888:80 -d -it bigdavidwong/camp:v1.0`
![img_2.png](img_2.png)

4. use nsenter to check ip config in container;

`docker ps`
![img_3.png](img_3.png)

`docker inspect b2b2b486ed50 | grep -i PID`
![img_4.png](img_4.png)

`nsenter -t 114887 -n ip addr`
![img_5.png](img_5.png)
