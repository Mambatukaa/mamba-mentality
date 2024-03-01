import { FolderIcon } from '@heroicons/react/24/outline';

export default function Header() {
  return (
    <header className="flex p-2 items-center col-span-2 bg-woodsmoke-950 text-white">
      <FolderIcon className={'h-7 w-7'} aria-hidden="true" />
    </header>
  );
}
