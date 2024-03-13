import { createSignal } from 'solid-js';
import { Button } from '~/components/ui/button';
import { Dialog, DialogContent, DialogTrigger } from '~/components/ui/dialog';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '~/components/ui/tabs';

import LoginForm from './login';
import RegisterForm from './register';

const LoginRegisterContainer = () => {
  const [openModal, setOpenModal] = createSignal(false);

  return (
    <Dialog open={openModal()} onOpenChange={setOpenModal}>
      {/* <DialogTrigger>
      
      </DialogTrigger> */}
      <Button onClick={() => setOpenModal(true)}>Login</Button>
      <DialogContent class="min-h-[475px]">
        <Tabs defaultValue="login" class="px-5 h-full">
          <TabsList class="grid w-full grid-cols-2">
            <TabsTrigger value="login">Login</TabsTrigger>
            <TabsTrigger value="register">Register</TabsTrigger>
          </TabsList>

          <TabsContent value="login" class="mt-8">
            <LoginForm />
          </TabsContent>

          <TabsContent value="register">
            <RegisterForm />
          </TabsContent>
        </Tabs>
      </DialogContent>
    </Dialog>
  );
};
export default LoginRegisterContainer;
