# create_project project_name [dir = .]
function create() {

	user_pwd=$(pwd)
	project_name=$1
	project_parent_dir=$2
	use_default_config=0

	default_permission="750"

	dir_build="build"
	dir_config="config"
	dir_doc="doc"
	dir_junk="junk"
	dir_logs="logs"
	dir_src="src"

	compile_path="/Users/Martin/2i/MPI/compile"
	run_path="/Users/Martin/2i/MPI/run"
	runtest_path="/Users/Martin/2i/MPI/runtest"
	username=`git config github.user`

	while [ -z $project_name ]
	do
		echo -n 'Project name : '	
		read project_name
	done

	while [ -z $project_parent_dir ]
	do
		echo -n 'Project parent dir : '
		read project_parent_dir
	done

	project_dir="$project_parent_dir/$project_name"


	echo "Creation of project '$project_name' in '$project_dir' ..."

	if [ ! -d "$project_dir" ]; then
		echo "'$project_dir' doesn't exist yet. Creation ..."
		mkdir $project_dir
		chmod 750 $project_dir
	fi

	cd $project_dir

	chmod 750 $project_dir

	echo -n "Do you want to use default config? [y/N] "
	read response
	case "$response" in
		[yY][eE][sS]|[yY]) echo "Using default config"
;;
*) echo "Enter directory name :"
echo -n "build (app builds) :"
read dir_build
echo -n "config (app configuration) :"
read dir_config
echo -n "doc (app documentation) :"
read dir_doc
echo -n "junk (app trash) :"
read dir_junk
echo -n "logs (app logs) :"
read dir_logs
echo -n "src (source code) :"
read dir_src
;;
esac

echo "Using config :"
echo "  $dir_build/"
echo "  $dir_config/"
echo "  $dir_doc/"
echo "  $dir_junk/"
echo "  $dir_logs/"
echo "  $dir_src/"

echo -n "Do you confirm? [y/N] "
read response
case "$response" in
	[yY][eE][sS]|[yY])
;;
*)
cd $user_pwd
return -1
;;
esac

mkdir $dir_build
mkdir $dir_config
mkdir $dir_doc
mkdir $dir_junk
mkdir $dir_logs
mkdir $dir_src

mkdir $dir_src/classes
mkdir $dir_src/classes/ClassName
cd $dir_src/classes/ClassName


touch ClassName.class.cpp
touch ClassName.class.hpp
touch ClassName.test.cpp
touch ClassName.test.dependencies

for i in *
do
	echo "// Project $project_name [duck-managed]" >> $i
	echo "// GitHub Repo : $project_name ()" >> $i
	echo "// $dir_src/classes/ClassName/$(basename $i)" >> $i
done

echo '
#ifndef CLASSNAME_CLASS_CPP
#define CLASSNAME_CLASS_CPP

#include "ClassName.class.hpp"

Class:Class() {
	//class constructor
}

#endif
' >> $dir_src/classes/ClassName/ClassName.class.cpp
echo '
// Project : 
#ifndef CLASSNAME_CLASS_HPP
#define CLASSNAME_CLASS_HPP


class Class
{
	public:
	Class(); //class constructor
	~Class();
	
};
#endif
' >> $dir_src/classes/ClassName/ClassName.class.hpp
mkdir $dir_src/config
mkdir $dir_src/tests
touch $dir_src/main.cpp

chmod 750 $dir_src/*


echo "Directory created."

cat $compile_path > compile
cat $run_path > run
cat $runtest_path > runtest

chmod $default_permission compile

echo "Created compilation, run and unit test script with permissions $default_permission"

echo "Starting creation of github repo"

echo ".DS_Store" > .gitignore

username=`git config github.user`
if [ "$username" = "" ]; then
	echo "Could not find git username, run 'git config --global github.user <username>'"
	invalid_credentials=1
fi

token=`git config github.token`
if [ "$token" = "" ]; then
	echo "Could not find token, run 'git config --global github.token <token>'"
	invalid_credentials=1
fi

if [ "$invalid_credentials" = "1" ]; then
	return 1
fi

echo -n "Creating Github repository '$repo_name' ..."
curl -u "$username:$token" https://api.github.com/user/repos -d '{"name":"'$repo_name'"}' > /dev/null 2>&1
echo " done."

commit_msg="Initial commit (duck-managed project)."

echo "Initializing the new repo ..."
git init
echo "Adding all files to commit ..."
git add *
echo "Commiting with message `$commit_msg`..."
git commit -m "$commit_msg"
echo -n "Pushing local code to remote ..."
git remote add origin https://github.com/$username/$repo_name.git > /dev/null 2>&1
git push -u origin master > /dev/null 2>&1
echo "Pushed successfully."


cd $user_pwd
}