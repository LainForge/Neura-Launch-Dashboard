import {
    Text,
     Card, 
     Grid, 
     Title,
     Subtitle,
     Button,
     Badge,
     TabGroup,
     TabList,
     Tab,
     TabPanels,
     TabPanel,
     Bold
} from "@tremor/react"
import UsersList from "../components/UsersList"

const models = [
    {
        title: "Speech Detection Model",
        subtitle: "/api.neuralaunch/kunxl.gg/predict",
        description: "This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.",
        status: "live"
    }, 

    {
        title: "Hate Speech Detection Model",
        subtitle: "/api.neuralaunch/grasgor/predict",
        description: "This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.",
        status: "not live"
    },

    {
        title: "Sentiment Analysis Model",
        subtitle: "/api.neuralaunch/akunxl.gg/predict",
        description: "This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.",
        status: "live"
    },

    {
        title: "Emotion Detection Model",
        subtitle: "/api.neuralaunch/akunxl.gg/predict",
        description: "This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.",
        status: "not live"
    }, 

    {
        title: "Speech Detection Model",
        subtitle: "/api.neuralaunch/kunxl.gg/predict",
        description: "This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.",
        status: "live"
    },

    {
        title: "Speech Detection Model",
        subtitle: "/api.neuralaunch/kunxl.gg/predict",
        description: "This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.",
        status: "live"
    },
]


export default function Dashboard(){
    return(
        <main className="w-screen min-h-screen px-5 bg-dark-tremor-background-muted">
            <SizedBox />
            <section className="flex flex-row justify-between" >
                <section>
                    <Bold className="text-white text-xl">Neura-Land</Bold>
                </section>
                <section className="flex flex-row justify-end">
                    <Button variant="light">Deploy</Button>
                    <VerticalSizedBox />
                    <Button variant="secondary">Costs</Button>
                    <VerticalSizedBox />
                    <Button variant="light">Usage</Button>
                    <VerticalSizedBox />
                    <Button variant="light">Users</Button>
                </section>
            </section>
            <SizedBox />
            <TabGroup>
                <TabList>
                    <Tab>Models</Tab>
                    <Tab>Users</Tab>
                </TabList>
                <SizedBox/>
                <TabPanels>
                    <TabPanel>
                        <DisplayGrid/>
                    </TabPanel>
                    <TabPanel>
                        <UsersList />
                    </TabPanel>
                </TabPanels>
            </TabGroup>
        </main>
    )
}

function DisplayGrid() {
    return(
        <main>
            <Grid numItems={3} className="gap-3">
                {models.map((model, index) => <DisplayCard title={model.title} subtitle={model.subtitle} description = {model.description} status = {model.status} key={index} />)}
            </Grid>
        </main>
    )
}

function DisplayCard(props: any) {
    return (
        <Card className="w-full"> 
            <section className="flex flex-row justify-between">
                <div>
                    <Title>{props.title}</Title>
                    <Subtitle>{props.subtitle}</Subtitle>
                </div>
                <div>{
                    props.status === "live" ? <Badge color="green">{props.status}</Badge> : <Badge color="red">{props.status}</Badge>
                    }
                </div>
            </section>
            <SizedBox/>
            <Text>{props.description}</Text>
            <SizedBox />
            <div className="w-full flex flex-between flex-row">
                <Button variant="secondary">Monitor</Button>
                <VerticalSizedBox />
                <Button variant="secondary">Modify</Button>
            </div>
        </Card>
    )
}

function SizedBox(){
    return(
        <main className="py-3"></main>
    )
}

function VerticalSizedBox() {
    return(
        <main className="px-3"></main>
    )
}

