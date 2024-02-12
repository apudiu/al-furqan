import { useLocation } from '@solidjs/router';
import Search from '~/components/ui/search';

const TopHeader = () => {
  const location = useLocation();
  console.log(location);
  return (
    <header class="h-20 border-b bg-muted px-5 py-2 w-full">
      <div class="flex justify-between items-center h-full">
        {/* <Show when={searchBarVisiblePaths.some((path) => location.pathname.startsWith(path))}>
          <Search />
        </Show> */}
        <Search />
        <p>User Dropdown</p>
      </div>
    </header>
  );
};
export default TopHeader;
