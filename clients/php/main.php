<?php

require __DIR__ .'/vendor/autoload.php';
require __DIR__ .'/../../pb/users/php/GPBMetadata/Users.php';
require __DIR__ .'/../../pb/users/php/Users/UserInfoServiceClient.php';
require __DIR__ .'/../../pb/users/php/Users/UserRequest.php';
require __DIR__ .'/../../pb/users/php/Users/UserResponse.php';

$opt = [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
];
$service = new \Users\UserInfoServiceClient('127.0.0.1:2333', $opt);

$req = new \Users\UserRequest();
$req->setName('wmh');
$resp = $service->GetUserInfo($req)->wait();
list($userInfo, $status) = $resp;

echo 'Name: ', $userInfo->getName(), "\n";
echo 'Age: ', $userInfo->getAge(), "\n";
$titles = $userInfo->getTitle();
foreach ($titles as $title) {
    echo "  --> ", $title, "\n";
}
