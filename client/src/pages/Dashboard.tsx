import {
    Text,
     Card, 
     Grid, 
     Title,
     Subtitle,
     Button,
     Badge
} from "@tremor/react"



export default function Dashboard(){
    return(
        <main>
           <DisplayGrid/> 
        </main>
    )
}

function DisplayGrid() {
    return(
        <main>
            <Grid numItems={3} className="gap-3">
                <DisplayCard />
                <DisplayCard />
                <DisplayCard />
            </Grid>
        </main>
    )
}

function DisplayCard() {
    return (
        <Card className="w-full"> 
            <section className="flex flex-row justify-between">
                <div>
                    <Title>Speech Detection Model</Title>
                    <Subtitle>/api.neuralaunch/kunxl.gg/predict</Subtitle>
                </div>
                <div>
                    <Badge color="red">live</Badge>
                </div>
            </section>
            <SizedBox/>
            <Text>This is model makde by noob people Lorem ipsum dolor sit amet consectetur adipisicing elit. Illum, aliquam.</Text>
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