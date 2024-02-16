import { Button } from '~/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '~/components/ui/card';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import { TextInput } from '~/components/ui/text-input';

const LoginForm = () => {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Login</CardTitle>
      </CardHeader>
      <CardContent class="space-y-2">
        <TextInput name="email" label="Email" error="" />
        <TextInput name="password" label="Password" error="" type="password" />
      </CardContent>
      <CardFooter>
        <Button type="submit">Login</Button>
      </CardFooter>
    </Card>
  );
};
export default LoginForm;
