import { ComponentProps, Show, splitProps } from 'solid-js';
import { Label } from './label';
import { Input } from './input';

type TextInputProps = {
  name: string;
  error: string;
  label?: string;
} & ComponentProps<'input'>;

export function TextInput(props: TextInputProps) {
  const [, inputProps] = splitProps(props, ['value', 'label', 'error']);

  return (
    <div>
      <Show when={props.label}>
        <Label>
          {props.label} {props.required && <span>*</span>}
        </Label>
      </Show>
      <Input
        {...inputProps}
        id={props.id}
        value={props.value || ''}
        aria-invalid={!!props.error}
        aria-errormessage={`${props.name}-error`}
      />

      {props.error && (
        <p id={`${props.name}-error`} class="text-sm font-medium text-destructive mt-[2px]">
          {props.error}
        </p>
      )}
    </div>
  );
}
