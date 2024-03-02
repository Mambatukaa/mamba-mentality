import { FolderIcon } from '@heroicons/react/24/outline';

export default function Header() {
  return (
    <header className="flex p-2 justify-between items-center col-span-2 bg-woodsmoke-950 text-white">
      <FolderIcon className={'h-7 w-7'} aria-hidden="true" />

      <div className="flex">
        <input
          className="block h-full border-0 py-0 pl-8 pr-0 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm"
          placeholder="Search"
          type="search"
          name="search"
        />
      </div>

      <button className="px-2 py-1 text-xs border-1 border-woodsmoke-500 rounded bg-woodsmoke-500">
        Store
      </button>
    </header>
  );
}
