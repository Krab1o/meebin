import { Button } from "~/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "~/components/ui/card"
import { Input } from "~/components/ui/input"
import { Label } from "~/components/ui/label"
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
} from "~/components/ui/tabs"
 
export function TabsDemo() {
  return (
    <Tabs defaultValue="account" className="w-[400px] mt-10">
      <TabsList className="grid w-full grid-cols-2">
        <TabsTrigger value="account">Аккаунт</TabsTrigger>
        <TabsTrigger value="password">Пароль</TabsTrigger>
      </TabsList>
      <TabsContent value="account">
        <Card>
          <CardHeader>
            <CardTitle>Аккаунт</CardTitle>
            <CardDescription>
                Внесите изменения в свой аккаунт здесь. Нажмите «Сохранить», когда закончите.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
            <div className="space-y-1">
              <Label htmlFor="name">Имя</Label>
              <Input id="name" defaultValue="Богдан Макконахи" />
            </div>
            <div className="space-y-1">
              <Label htmlFor="username">Логин</Label>
              <Input id="username" defaultValue="@krab1o" />
            </div>
          </CardContent>
          <CardFooter>
            <Button>Сохранить изменения</Button>
          </CardFooter>
        </Card>
      </TabsContent>
      <TabsContent value="password">
        <Card>
          <CardHeader>
            <CardTitle>Пароль</CardTitle>
            <CardDescription>
              Измените свой пароль. После сохранения вы сможете войти в систему.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-2">
            <div className="space-y-1">
              <Label htmlFor="current">Текущий пароль</Label>
              <Input id="current" type="password" />
            </div>
            <div className="space-y-1">
              <Label htmlFor="new">Новый пароль</Label>
              <Input id="new" type="password" />
            </div>
          </CardContent>
          <CardFooter>
            <Button>Сохранить пароль</Button>
          </CardFooter>
        </Card>
      </TabsContent>
    </Tabs>
  )
}