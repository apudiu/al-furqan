import { Button } from '~/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '~/components/ui/card';
import { TextInput } from '~/components/ui/text-input';

const RegisterForm = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Create Your Account</CardTitle>
      </CardHeader>
      <CardContent class="space-y-2">
        <TextInput name="email" label="Email" error="" />
        <TextInput name="password" label="Password" error="" type="password" />
        <TextInput name="name" label="Name" error="" />
      </CardContent>
      <CardFooter>
        <Button type="submit">Create Account</Button>
      </CardFooter>
    </Card>
  );
};
export default RegisterForm;
