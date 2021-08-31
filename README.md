# Software-Tech-Test

One of DroneShield's missions is to provide the best Counterdrone defence in an emerging industry.
This challenge involves building a simulator for Counter UAS involving the detection of drones.
Please spend no more than 6 hours on this challenge as we want you to have a life :). If you do more tasks
then we expect less detail in the individual tasks given the time constraints. This should be a fun exercise.

Please provde an explanation with what could be done to improve your work if given more time (bullet points acceptable). Feel free to use any libraries and discuss advanced concepts in the readme because our challenge does not showcase your vast knowledge. :D

## Instructions to run the sample
```
# The sample starts at Sydney airport and flies diagonally indefinitely, feel free to add an upperbound or reverse the direction.
npm i
npm start
# Websocket endpoint, use your favourite WS client
ws://localhost:8080
```

## Expectations for Engineers
- <b>Frontend Engineers</b> (Frontend Task)
- <b>Backend Engineers</b>  (Backend Task)
- <b>Fullstack Engineers</b>  (Frontend and Backend Tasks [Choice of <b>one</b> or simple versions of <b>both</b>])
- <b>Lead</b> or <b>Polyglot Engineers</b>  (Frontend, Backend and Infrastructure Tasks [Choice of <b>two</b> or simple versions of all <b>three</b>])

## Frontend Task
- Technology Preference: VueJs or ReactJs
- Objective: Make a page that displays the drone movement through the websocket in realtime such as a map or summary page. Time is limited so pick the easier of the two if you are completing multiple tasks. If only completing this taks please display both, feel free to add more mock detail to the websocket.
- Design System choice (A brief explanation of your chosen UI library and Why) (<b>Required</b>)
- <b>Junior Engineers:</b> Using known libraries without customisation is acceptable.
- <b>Midtier Engineers:</b> More emphasis will be given to CSS and custom component design (not necessarily entirely from scratch but able to show customisation of base components from a UI library like Material UI). Your work will demonstrate knowledge of the state management lifecycle.
- <b>Senior Engineers:</b> Data flow and rendering efficiency (performance) will be prioritised. Your work will demonstrate such concepts such as code splitting (lazy loading) and where to connect state at parent or child components. If you feel your design implementation doesn't need to contain these feel free to discuss these topics indepth in the readme and how it prevents excessive renders and/or improves performance.

## Backend Task
- Technology Preference: NodeJs, GoLang, Ruby or Python
- Objective: Make two microservices that interface with each through redis publish and subscribe. One microservice that publishes to redis the coordinates of the drone and another microservice that subscribes to the event and pushes to a websocket (similar to the sample). Additionally make a selectable simulated flying pattern of the drone such as "Figure 8", "Circle" or "Zigzag". More advanced and intelligent flight paths will be given additional brownie points.
- <b>Junior Engineers:</b> Single microservice as a monolith is acceptible (no pubsub required just websocket and custom drone path).
- <b>Midtier Engineers:</b> Substituting Redis pubsub with Rest API for intra microservice communication is acceptable.
- <b>Senior Engineers:</b> Other queue like technologies instead of redis such as RabbitMQ and Kafka are acceptable. As the challenge is fairly simple feel free to discuss microservice architecture to improve performance, techniques for seperation of concerns, clean architecture and/or domain driven design.
- Architecture Diagram (<b>Required</b>)

## Infrastructure Task
- Technology Preference: Docker-compose or Kubernetes
- Objective: Dockerize the above with docker build files and provide a runnable docker-compose solution. If feeling creative provide kubernetes/helm templates to run on minikube.
- Architecture Diagram (<b>Required</b>)