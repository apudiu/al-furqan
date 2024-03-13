import { A } from '@solidjs/router';
import SideNav from './nav';
import Logo from '~/components/ui/logo';
import { routePaths } from '~/lib/routes';

const Sidebar = () => {
  return (
    <aside class="bg-primary h-screen max-h-screen overflow-y-auto overflow-x-hidden p-5 pb-20 flex flex-col gap-y-8">
      <A href={routePaths.home}>
        <Logo />
      </A>
      <SideNav />
    </aside>
  );
};
export default Sidebar;
