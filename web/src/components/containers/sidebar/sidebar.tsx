import { A } from '@solidjs/router';
import SideNav from './nav';
import Logo from '~/components/ui/logo';
import { routePaths } from '~/lib/routes';

const Sidebar = () => {
  return (
    <aside class="bg-primary h-screen max-h-screen overflow-y-auto overflow-x-hidden p-5 space-y-8">
      <A href={routePaths.home} class="h-10">
        <Logo />
      </A>
      <div class="flex flex-col justify-between h-full max-h-[calc(100vh_-_100px)]">
        <SideNav />
        <footer class="text-white">Footer</footer>
      </div>
    </aside>
  );
};
export default Sidebar;
