import { createSignal } from 'solid-js';
import { Button } from '~/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '~/components/ui/card';
import { Dialog, DialogContent, DialogTrigger } from '~/components/ui/dialog';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '~/components/ui/tabs';
import { TextInput } from '~/components/ui/text-input';

const LoginRegisterContainer = () => {
  const [openModal, setOpenModal] = createSignal(false);

  return (
    <Dialog open={openModal()} onOpenChange={setOpenModal}>
      <DialogTrigger>
        <Button>Login</Button>
      </DialogTrigger>
      <DialogContent class="min-h-[475px]">
        <Tabs defaultValue="login" class="px-5 h-full">
          <TabsList class="grid w-full grid-cols-2">
            <TabsTrigger value="login">Login</TabsTrigger>
            <TabsTrigger value="register">Register</TabsTrigger>
          </TabsList>

          <TabsContent value="login" class="mt-8">
            <Card>
              <CardHeader>
                <CardTitle>Login</CardTitle>
              </CardHeader>
              <CardContent class="space-y-2">
                <div class="space-y-1">
                  {/* <Label for="email">Email</Label>
                  <Input id="email" /> */}
                  <TextInput name="email" label="Email" error="" />
                </div>
                <div class="space-y-1">
                  <Label for="password">password</Label>
                  <Input id="password" type="password" />
                </div>
              </CardContent>
              <CardFooter>
                <Button>Login</Button>
              </CardFooter>
            </Card>
          </TabsContent>

          <TabsContent value="register">
            <Card>
              <CardHeader>
                <CardTitle>Create Your Account</CardTitle>
              </CardHeader>
              <CardContent class="space-y-2">
                <div class="space-y-1">
                  {/* <Label for="name">Name</Label> */}
                  {/* <Input id="name" /> */}
                  <TextInput name="name" label="Name" error="" />
                </div>

                <div class="space-y-1">
                  <Label for="email">Email</Label>
                  <Input id="email" />
                </div>
                <div class="space-y-1">
                  <Label for="password">password</Label>
                  <Input id="password" type="password" />
                </div>
              </CardContent>
              <CardFooter>
                <Button>Create Account</Button>
              </CardFooter>
            </Card>
          </TabsContent>
        </Tabs>
      </DialogContent>
    </Dialog>
  );
};
export default LoginRegisterContainer;
