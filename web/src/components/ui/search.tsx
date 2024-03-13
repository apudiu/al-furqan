import { useNavigate } from '@solidjs/router';
import { Input } from './input';

import { SubmitHandler, createForm } from '@modular-forms/solid';
import { routePaths } from '~/lib/routes';

type SearchForm = {
  query: string;
};

const Search = () => {
  const navigate = useNavigate();

  const [_, { Form, Field, FieldArray }] = createForm<SearchForm>();

  const handleSubmit: SubmitHandler<SearchForm> = (values, event) => {
    const path = routePaths.search + '/' + values.query;
    navigate(path);
  };

  return (
    <Form onSubmit={handleSubmit} class="w-full max-w-md">
      <Field name="query">
        {(field, props) => (
          <Input {...props} id={field.name} value={field.value} placeholder="Search" />
        )}
      </Field>
    </Form>
  );
};
export default Search;
