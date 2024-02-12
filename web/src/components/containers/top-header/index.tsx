import Search from '~/components/ui/search';
import LoginRegisterContainer from '../auth/login-register-container';

const TopHeader = () => {
  return (
    <header class="h-20 border-b bg-muted px-5 py-2 w-full">
      <div class="flex justify-between items-center h-full">
        <Search />

        <div class="ml-auto">
          <LoginRegisterContainer />
        </div>
      </div>
    </header>
  );
};
export default TopHeader;
