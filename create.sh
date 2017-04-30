# create_project project_name [dir = .]
project_name=$1
project_parent_dir=$2
project_dir=""$project_dir/$project_name""


cd $project_parent_dir



echo "Creation of project $project_name in $project_parent_dir..."

if [ ! -d "$project_dir" ]; then
	echo "$project_dir doesn't exist yet. Creation ..."
	mkdir $project_dir
	chmod 750 $project_dir
fi

chmod 750 $project_dir

echo "Do you wish to install this program?"
select yn in "Yes" "No"; do
    case $yn in
        Yes ) make install; break;;
        No ) exit;;
    esac
done

$dir_build = "build/"

date=
while [ -z $date ]
do
    echo -n 'Date? '
    read date
done